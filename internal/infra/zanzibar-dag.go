package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
)

type ZanzibarDagRepository struct {
	Url string
}

func NewZanzibarDagRepository() *ZanzibarDagRepository {
	host := os.Getenv("ZANZIBAR_DAG_HOST")
	port := os.Getenv("ZANZIBAR_DAG_PORT")
	return &ZanzibarDagRepository{
		Url: fmt.Sprintf("http://%s:%s/relation", host, port),
	}
}

func (r *ZanzibarDagRepository) GetAll() ([]zanzibardagdom.Relation, error) {
	resp, err := http.Get(r.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var relations []zanzibardagdom.Relation
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations, nil
}

func (r *ZanzibarDagRepository) Query(relation zanzibardagdom.Relation) ([]zanzibardagdom.Relation, error) {

	resp, err := http.Get(r.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var relations []zanzibardagdom.Relation
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations, nil
}

func (r *ZanzibarDagRepository) Create(relation zanzibardagdom.Relation) error {
	payload, err := json.Marshal(relation)
	if err != nil {
		return err
	}

	resp, err := http.Post(r.Url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (r *ZanzibarDagRepository) Delete(relation zanzibardagdom.Relation) error {
	payload, err := json.Marshal(relation)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", r.Url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (r *ZanzibarDagRepository) GetAllNamespaces() ([]string, error) {
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the GET request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body to extract namespaces
	var namespacesResponse struct {
		Data []string `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&namespacesResponse); err != nil {
		return nil, err
	}

	return namespacesResponse.Data, nil
}

func (r *ZanzibarDagRepository) Check(from zanzibardagdom.Node, to zanzibardagdom.Node) (bool, error) {

	payload := map[string]interface{}{
		"subject_namespace": from.Namespace,
		"subject_name":      from.Name,
		"subject_relation":  from.Relation,
		"object_namespace":  to.Namespace,
		"object_name":       to.Name,
		"relation":          to.Relation,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", r.Url+"/relation/check", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return true, nil
}

func (r *ZanzibarDagRepository) GetShortestPath(from zanzibardagdom.Node, to zanzibardagdom.Node) ([]zanzibardagdom.Relation, error) {

	payload := map[string]interface{}{
		"subject_namespace": from.Namespace,
		"subject_name":      from.Name,
		"subject_relation":  from.Relation,
		"object_namespace":  to.Namespace,
		"object_name":       to.Name,
		"relation":          to.Relation,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", r.Url+"/relation/get-shortest-path", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var relations []zanzibardagdom.Relation
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations, nil
}

func (r *ZanzibarDagRepository) GetAllPaths(from zanzibardagdom.Node, to zanzibardagdom.Node) ([][]zanzibardagdom.Relation, error) {

	payload := map[string]interface{}{
		"subject_namespace": from.Namespace,
		"subject_name":      from.Name,
		"subject_relation":  from.Relation,
		"object_namespace":  to.Namespace,
		"object_name":       to.Name,
		"relation":          to.Relation,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", r.Url+"/relation/get-all-paths", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var paths [][]zanzibardagdom.Relation
	if err := json.NewDecoder(resp.Body).Decode(&paths); err != nil {
		return nil, err
	}

	return paths, nil
}

func (r *ZanzibarDagRepository) GetAllObjectRelations(subject zanzibardagdom.Node) ([]zanzibardagdom.Relation, error) {

	payload := map[string]interface{}{
		"namespace": subject.Namespace,
		"name":      subject.Name,
		"relation":  subject.Relation,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", r.Url+"/relation/get-all-object-relations", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var relations []zanzibardagdom.Relation
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations, nil
}

func (r *ZanzibarDagRepository) GetAllSubjectRelations(object zanzibardagdom.Node) ([]zanzibardagdom.Relation, error) {

	payload := map[string]interface{}{
		"namespace": object.Namespace,
		"name":      object.Name,
		"relation":  object.Relation,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", r.Url+"/relation/get-all-subject-relations", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var relations []zanzibardagdom.Relation
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations, nil
}

func (r *ZanzibarDagRepository) ClearAllRelations() error {

	req, err := http.NewRequest("POST", r.Url+"/relation/clear-all-relations", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
